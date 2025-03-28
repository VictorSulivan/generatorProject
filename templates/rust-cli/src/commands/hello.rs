use clap::Parser;
use crate::error::Result;

#[derive(Parser, Debug)]
pub struct HelloCommand {
    /// Nom à saluer
    #[arg(short, long, default_value = "World")]
    name: String,
}

impl HelloCommand {
    pub fn execute(&self) -> Result<()> {
        println!("Hello, {}!", self.name);
        Ok(())
    }
} 