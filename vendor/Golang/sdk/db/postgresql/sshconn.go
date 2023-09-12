package postgresql

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Open(s string) (_ driver.Conn, err error) {
	return pq.DialOpen(self, s)
}

func (self *ViaSSHDialer) Dial(network, address string) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func (self *ViaSSHDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func main() {

	sshHost := "example.com" // SSH Server Hostname/IP
	sshPort := 22            // SSH Port
	sshUser := "ssh-user"    // SSH Username
	sshPass := "ssh-pass"    // Empty string for no password
	dbUser := "user"         // DB username
	dbPass := "password"     // DB Password
	dbHost := "localhost"    // DB Hostname/IP
	dbName := "database"     // Database name

	var agentClient agent.Agent
	// Establish a connection to the local ssh-agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()

		// Create a new instance of the ssh agent
		agentClient = agent.NewClient(conn)
	}

	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{},
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}
	// When there's a non empty password add the password AuthMethod
	if sshPass != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return sshPass, nil
		}))
	}

	// Connect to the SSH Server
	if sshcon, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshHost, sshPort), sshConfig); err == nil {
		defer sshcon.Close()

		// Now we register the ViaSSHDialer with the ssh connection as a parameter
		sql.Register("postgres+ssh", &ViaSSHDialer{sshcon})

		// And now we can use our new driver with the regular postgres connection string tunneled through the SSH connection
		if db, err := sql.Open("postgres+ssh", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbName)); err == nil {

			fmt.Printf("Successfully connected to the db\n")

			if rows, err := db.Query("SELECT id, name FROM table ORDER BY id"); err == nil {
				for rows.Next() {
					var id int64
					var name string
					rows.Scan(&id, &name)
					fmt.Printf("ID: %d  Name: %s\n", id, name)
				}
				rows.Close()
			}

			db.Close()

		} else {

			fmt.Printf("Failed to connect to the db: %s\n", err.Error())
		}

	}
}
