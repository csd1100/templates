use anyhow::{anyhow, Result};

pub fn add(num1: i32, num2: i32) -> Result<i32> {
    return Ok(num1 + num2);
}

pub fn div(num1: i32, num2: i32) -> Result<i32> {
    if num2 == 0 {
        return Err(anyhow!("Cannot divide by zero"));
    }

    return Ok(num1 / num2);
}

#[cfg(test)]
mod test {
    use anyhow::Result;

    use crate::math::math::{add, div};

    #[test]
    fn add_2_numbers() -> Result<()> {
        assert_eq!(2, add(1, 1)?);
        return Ok(());
    }

    #[test]
    fn div_2_numbers() -> Result<()> {
        assert_eq!(2, div(4, 2)?);
        return Ok(());
    }

    #[test]
    fn div_by_zero() -> Result<()> {
        let expected = String::from("Cannot divide by zero");
        let actual = div(100, 0).unwrap_err().to_string();
        assert_eq!(expected, actual);
        return Ok(());
    }
}
