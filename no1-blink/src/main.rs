#![no_std]
#![no_main]

use panic_halt as _;

#[arduino_hal::entry]
fn main() -> ! {
    let dp = arduino_hal::Peripherals::take().unwrap();
    let pins = arduino_hal::pins!(dp);

    let mut led = pins.d13.into_output();
    led.set_high();

    loop {
        led.toggle();
        arduino_hal::delay_ms(1000);
        led.toggle();
        arduino_hal::delay_ms(2000);
        led.toggle();
        arduino_hal::delay_ms(3000);
        led.toggle();
        arduino_hal::delay_ms(5000);
    }
}