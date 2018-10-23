use std::{i8, i16, i32, i64, u8, u16, u32, u64, isize, usize, f32, f64};

use std::io::stdin;

fn main() {
    println!("Hello World!");

    let num = 10;

    let mut age: i32 = 23;
    let let_x: char = 'x';
    let is_it_true: bool = true;
    println!("It is {0} that {1} is {0}", is_it_true, let_x);

    println!("I am {} years old", age);

    let (f_name, l_name) = ("Jesse", "Barton");
    println!("My name is {} {}", f_name, l_name);
    println!("Bin: {:b} Hex: {:x} Octa: {:o}", num, num, num);
}