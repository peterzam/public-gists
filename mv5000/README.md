# mv5000.sh

### Move the first 5000 files into _5000 temporary folder

file usage:

```bash
bash mv5000.sh pictures p0
```

```bash
mkdir $2;ls $1 | head -5000 | xargs -I{} mv ./$1/{} ./$2
```
