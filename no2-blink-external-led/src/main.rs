#![no_std]
#![no_main]

use panic_halt as _;

#[arduino_hal::entry]
fn main() -> ! {
    let dp = arduino_hal::Peripherals::take().unwrap();
    let pins = arduino_hal::pins!(dp);

    let mut external_led = pins.d12.into_output();
    external_led.set_high();

    let mut external_led_2 = pins.d8.into_output();
    external_led_2.set_low();

    loop {
        external_led.toggle();
        external_led_2.toggle();
        arduino_hal::delay_ms(1000);

        external_led.toggle();
        external_led_2.toggle();
        arduino_hal::delay_ms(1000);

        external_led.toggle();
        external_led_2.toggle();
        arduino_hal::delay_ms(1000);
    }
}
