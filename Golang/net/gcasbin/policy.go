package gcasbin

// ACL: Access Control List
// rbac: Role Based Access Control
// abac: Attribute-based access control

const PolicyACL = `
p, alice, data1, read
p, bob, data2, write
`

const PolicyACLWithSuperuser = `
p, alice, data1, read
p, bob, data2, write
`

const PolicyACLWithoutResource = `
p, alice, read
p, bob, write
`

const PolicyACLWithoutUser = `
p, data1, read
p, data2, write
`

const PolicyRbac = `
p, alice, data1, read
p, bob, data2, write
p, data2_admin, data2, read
p, data2_admin, data2, write

g, alice, data2_admin
`

const PolicyRbacResource = `
p, alice, data1, read
p, bob, data2, write
p, data_group_admin, data_group, write

g, alice, data_group_admin
g2, data1, data_group
g2, data2, data_group
`

const PolicyRbacWithDomain = `
p, admin, domain1, data1, read
p, admin, domain1, data1, write
p, admin, domain2, data2, read
p, admin, domain2, data2, write

g, alice, admin, domain1
g, bob, admin, domain2
`

const PolicyRbacWithPattern = `
p, pen_admin, data1, GET
g, /book/:id, pen_admin
`

const PolicyRbacWithPattern2 = `
p, book_group, domain1, data1, read
p, book_group, domain2, data2, write

g, /book/:id, book_group, *
`

const PolicyRbacWithDeny = `
p, alice, data1, read, allow
p, bob, data2, write, allow
p, data2_admin, data2, read, allow
p, data2_admin, data2, write, allow
p, alice, data2, write, deny

g, alice, data2_admin
`

const PolicyAbac = `

`

const PolicyAbacWithPolicy = `
p, r.sub.Age > 18 && r.sub.Age < 60, /data1, read
`

const PolicyRestfulMatch = `
p, alice, /alice_data/*, GET
p, alice, /alice_data/resource1, POST

p, bob, /alice_data/resource2, GET
p, bob, /bob_data/*, POST

p, cathy, /cathy_data, (GET)|(POST)
`

const PolicyRestfulMatch2 = `
p, alice, /alice_data/:resource, GET
p, alice, /alice_data2/:id/using/:resId, GET
`

const PolicyIPMatch = `
p, 192.168.2.0/24, data1, read
p, 10.0.0.0/16, data2, write
`

const PolicyPriority = `
p, alice, data1, read, allow
p, data1_deny_group, data1, read, deny
p, data1_deny_group, data1, write, deny
p, alice, data1, write, allow

g, alice, data1_deny_group

p, data2_allow_group, data2, read, allow
p, bob, data2, read, deny
p, bob, data2, write, deny

g, bob, data2_allow_group
`
