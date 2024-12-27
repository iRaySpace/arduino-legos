#![no_std]
#![no_main]

use panic_halt as _;

#[arduino_hal::entry]
fn main() -> ! {
    let dp = arduino_hal::Peripherals::take().unwrap();
    let pins = arduino_hal::pins!(dp);

    let ir = pins.d3.into_floating_input();
    let mut led = pins.d9.into_output();

    loop {
        if ir.is_low() {
            led.set_high();
        } else {
            led.set_low();
        }
        arduino_hal::delay_ms(300);
    }
}
