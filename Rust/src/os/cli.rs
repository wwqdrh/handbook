
use clap::Parser;

use anyhow::Result;

#[derive(Parser, Debug)]
#[clap(
    author,
    version,
    about,
    long_about = "A TUI toolkit to view HelloGitHub"
)]
pub struct Args {
    #[clap(short, long, help = "配置文件路径")]
    pub path: Option<String>,

    #[clap(short, long, help = "是否显示帮助")]
    pub show_help: bool,

    #[clap(long, help = "显示内置样式列表")]
    pub show_themes: bool,
}

#[allow(dead_code)]
#[derive(Debug, Default, Clone)]
pub struct Config {
    pub config_path: String,
    pub show_help: bool,
    pub show_themes: bool,
}

impl From<Args> for Config {
    fn from(args: Args) -> Self {
        #[cfg(not(target_os = "windows"))]
        let home = env!("HOME").to_string();
        #[cfg(target_os = "windows")]
        let home = env!("HOMEPATH").to_string();

        let config_path = args.path.unwrap_or(home);
        Config {
            config_path,
            show_help: args.show_help,
            show_themes: args.show_themes,
        }
    }
}

pub fn parse_args() -> Result<Config> {
    let args = Args::parse();
    Ok(Config::from(args))
}