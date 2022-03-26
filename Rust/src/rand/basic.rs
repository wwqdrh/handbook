use rand::Rng;

pub fn randu32() -> u32{
    let mut rng = rand::thread_rng();

    rng.gen::<u32>()
}

#[cfg(test)]
mod basic_tests {
    use super::*;

    #[test]
    fn test_randu32() {
        println!("restult: {}", randu32())
    }
}