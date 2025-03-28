use std::path::PathBuf;
use crate::error::Result;
use crate::TemplateType;
use walkdir::WalkDir;
use handlebars::Handlebars;
use std::fs;

pub struct Generator {
    templates_dir: PathBuf,
    handlebars: Handlebars<'static>,
}

impl Generator {
    pub fn new() -> Result<Self> {
        let templates_dir = PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("templates");
        let mut handlebars = Handlebars::new();
        handlebars.set_strict_mode(true);
        
        Ok(Self {
            templates_dir,
            handlebars,
        })
    }

    pub fn generate(&self, template: TemplateType, name: &str, output: Option<PathBuf>) -> Result<()> {
        let output_dir = output.unwrap_or_else(|| PathBuf::from(name));
        
        // Créer le répertoire de sortie
        fs::create_dir_all(&output_dir)?;

        // Copier et traiter les fichiers du template
        let template_dir = self.templates_dir.join(template.as_str());
        for entry in WalkDir::new(&template_dir) {
            let entry = entry?;
            let path = entry.path();
            let relative_path = path.strip_prefix(&template_dir)?;
            let target_path = output_dir.join(relative_path);

            if path.is_file() {
                // Créer les répertoires parents si nécessaire
                if let Some(parent) = target_path.parent() {
                    fs::create_dir_all(parent)?;
                }

                // Lire le contenu du fichier
                let content = fs::read_to_string(path)?;

                // Remplacer les variables si c'est un template
                let processed_content = if path.extension().map_or(false, |ext| ext == "hbs") {
                    let mut data = serde_json::json!({
                        "ProjectName": name,
                    });
                    self.handlebars.render_template(&content, &data)?
                } else {
                    content
                };

                // Écrire le fichier de sortie
                let output_path = if path.extension().map_or(false, |ext| ext == "hbs") {
                    target_path.with_extension("")
                } else {
                    target_path
                };
                fs::write(output_path, processed_content)?;
            }
        }

        println!("Template généré avec succès dans {}", output_dir.display());
        Ok(())
    }
} 