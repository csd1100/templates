use ___projectName___::math::math::{add, div};
use anyhow::Result;

fn main() -> Result<()> {
    println!("Hello World");
    let a = 100;
    let b = 120;
    let sum = add(a, b)?;
    println!("{} + {} = {}", a, b, sum);
    let a = 400;
    let b = 120;
    let div = div(a, b)?;
    println!("{} / {} = {}", a, b, div);
    return Ok(());
}
