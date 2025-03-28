use clap::Parser;
use std::path::PathBuf;
use crate::error::Result;
use std::fs;

#[derive(Parser, Debug)]
pub struct FileCommand {
    /// Chemin du fichier
    #[arg(short, long)]
    path: PathBuf,

    /// Action à effectuer (read, write, delete)
    #[arg(short, long)]
    action: FileAction,

    /// Contenu à écrire (pour l'action write)
    #[arg(short, long)]
    content: Option<String>,
}

#[derive(Parser, Debug, Clone, Copy)]
pub enum FileAction {
    /// Lire le fichier
    Read,
    /// Écrire dans le fichier
    Write,
    /// Supprimer le fichier
    Delete,
}

impl FileCommand {
    pub fn execute(&self) -> Result<()> {
        match self.action {
            FileAction::Read => {
                let content = fs::read_to_string(&self.path)?;
                println!("Contenu du fichier:\n{}", content);
            }
            FileAction::Write => {
                let content = self.content.as_ref()
                    .ok_or_else(|| crate::error::AppError::Config("Contenu requis pour l'écriture".to_string()))?;
                fs::write(&self.path, content)?;
                println!("Fichier écrit avec succès");
            }
            FileAction::Delete => {
                fs::remove_file(&self.path)?;
                println!("Fichier supprimé avec succès");
            }
        }
        Ok(())
    }
} 