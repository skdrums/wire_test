for d in `find /Volumes/UNTITLED/images/ -type d`; do echo $d,`ls "$d" | wc -l`; done
