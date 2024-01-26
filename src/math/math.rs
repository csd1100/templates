pub fn add(num1: i32, num2: i32) -> i32 {
    return num1 + num2;
}

#[cfg(test)]
mod test {
    use crate::math::math::add;

    #[test]
    fn add_2_numbers() {
        assert_eq!(2, add(1, 1));
    }
}
