use clap::Parser;
use crate::error::Result;

#[derive(Parser, Debug)]
pub struct CalcCommand {
    /// Premier nombre
    #[arg(short, long)]
    a: f64,

    /// Deuxième nombre
    #[arg(short, long)]
    b: f64,

    /// Opération à effectuer (add, sub, mul, div)
    #[arg(short, long)]
    operation: Operation,
}

#[derive(Parser, Debug, Clone, Copy)]
pub enum Operation {
    /// Addition
    Add,
    /// Soustraction
    Sub,
    /// Multiplication
    Mul,
    /// Division
    Div,
}

impl CalcCommand {
    pub fn execute(&self) -> Result<()> {
        let result = match self.operation {
            Operation::Add => self.a + self.b,
            Operation::Sub => self.a - self.b,
            Operation::Mul => self.a * self.b,
            Operation::Div => {
                if self.b == 0.0 {
                    return Err(crate::error::AppError::Config("Division par zéro".to_string()));
                }
                self.a / self.b
            }
        };

        println!("Résultat: {}", result);
        Ok(())
    }
} 