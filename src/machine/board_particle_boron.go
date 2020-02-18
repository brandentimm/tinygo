// +build particle_boron

package machine

const HasLowFrequencyCrystal = true

// More info: https://docs.particle.io/datasheets/cellular/boron-datasheet/
// Board diagram: https://docs.particle.io/assets/images/boron/boron-block-diagram.png

// LEDs
const (
	LED       Pin = 44
	LED_GREEN Pin = 14
	LED_RED   Pin = 13
	LED_BLUE  Pin = 15
)

// UART pins
const (
	UART_TX_PIN Pin = 6
	UART_RX_PIN Pin = 8
)

// I2C pins
const (
	SDA_PIN Pin = 26
	SCL_PIN Pin = 27

	// Internal I2C with MAX17043 and BQ24195 chips on it
	SDA1_PIN Pin = 24
	SCL1_PIN Pin = 41
	INT1_PIN Pin = 5
)

// SPI pins
const (
	SPI0_SCK_PIN  Pin = 47
	SPI0_MOSI_PIN Pin = 45
	SPI0_MISO_PIN Pin = 46
)

// Internal 4MB SPI Flash
const (
	SPI1_SCK_PIN  Pin = 19
	SPI1_MOSI_PIN Pin = 20
	SPI1_MISO_PIN Pin = 21
	SPI1_CS_PIN   Pin = 17
	SPI1_WP_PIN   Pin = 22
	SPI1_HOLD_PIN Pin = 23
)

// u-blox coprocessor
const (
	UBLOX_TXD_PIN      Pin = 37
	UBLOX_RXD_PIN      Pin = 36
	UBLOX_CTS_PIN      Pin = 38
	UBLOX_RTS_PIN      Pin = 39
	UBLOX_RESET_PIN    Pin = 16
	UBLOX_POWER_ON_PIN Pin = 24
	UBLOX_BUFF_EN_PIN  Pin = 25
	UBLOX_VINT_PIN     Pin = 2
)

// Other periferals
const (
	MODE_BUTTON_PIN   Pin = 11
	CHARGE_STATUS_PIN Pin = 41
	LIPO_VOLTAGE_PIN  Pin = 5
	PCB_ANTENNA_PIN   Pin = 2
	EXTERNAL_UFL_PIN  Pin = 25
	NFC1_PIN          Pin = 9
	NFC2_PIN          Pin = 10
)