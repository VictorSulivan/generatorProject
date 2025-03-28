use clap::Parser;
use tracing::info;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    /// Commande à exécuter
    #[arg(short, long)]
    command: String,
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    // Initialisation du logging
    tracing_subscriber::fmt::init();
    info!("Starting {{.ProjectName}}");

    // Parsing des arguments
    let args = Args::parse();

    // Exécution de la commande
    match args.command.as_str() {
        "hello" => println!("Hello, World!"),
        _ => println!("Unknown command: {}", args.command),
    }

    Ok(())
} 