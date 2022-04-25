# Usage

## Initial

```
rhi init
```

## ASCII Art

```
rhi ascart <image_path>
rhi ascart -x 100 -y 50 <image_path>
```

## Download

### download image

```
rhi dl img -u <url> -s <html5 selector> -o <outdir> 

# limit
rhi dl img -u <url> -s <html5 selector> -o <outdir> -l 10

# concurrency
rhi dl img -u <url> -s <html5 selector> -o <outdir> -l 10 -c 2
```

## Random

```
rhi rand [length]

# all Uppercase
rhi rand -U

# allow uppercase
rhi rand -u

# allow symbol
rhi rand -s
```
