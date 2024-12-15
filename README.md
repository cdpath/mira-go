# Mira-go

A Go implementation of [mira-js](https://github.com/ipodnerd3019/mira-js)

## Installation

```bash
go install github.com/cdpath/mira-go/cmd/mira@latest
```

## Usage

```bash
# Refresh the screen
mira refresh

# Apply settings
mira settings --speed 3 --contrast 11

# Enable anti-shake
mira antishake

# Get help
mira --help
```

### Available Commands

- `refresh`: Refresh the screen
- `antishake`: Enable anti-shake automatically
- `settings`: Apply various display settings

### Settings Options

- `--speed`: The refresh speed (1-7)
- `--contrast`: The contrast (0-15)
- `--refresh-mode`: The refresh mode (a2, direct, gray)
- `--dither-mode`: The dither mode (0-3)
- `--black-filter`: The black filter level (0-254)
- `--white-filter`: The white filter level (0-254)
- `--cold-light`: The cold backlight level (0-254)
- `--warm-light`: The warm backlight level (0-254)

## Modes

Modes are just combinations of settings flags.

### Speed
```bash
mira settings --refresh-mode a2 --contrast 8 --speed 7 --dither-mode 0 --white-filter 0 --black-filter 0
```

### Text
```bash
mira settings --refresh-mode a2 --contrast 7 --speed 6 --dither-mode 1 --white-filter 0 --black-filter 0
```

### Image
```bash
mira settings --refresh-mode direct --contrast 7 --speed 5 --dither-mode 0 --white-filter 0 --black-filter 0
```

### Video
```bash
mira settings --refresh-mode a2 --contrast 7 --speed 6 --dither-mode 2 --white-filter 10 --black-filter 0
```

### Read
```bash
mira settings --refresh-mode direct --contrast 7 --speed 5 --dither-mode 3 --white-filter 12 --black-filter 10
```
