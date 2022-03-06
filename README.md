# imgaug

Image augmentation from the command line.

I wrote this to help synthesize new images for use in computer vision machine learning.

I realize there's plenty of python packages that can augment images. They are probably more advanced.

However, I wanted a simple command line tool I could run immediately without needing to set up a python environment,
download packages, etc. 

## Installation

You will need Go installed.

If you use [asdf](https://github.com/asdf-vm/asdf) (highly recommended):

```bash
asdf install
```

Ensure `~/go/bin` in your `$PATH`.

```bash
make 
```

## Usage

```bash
imgaug --help
```