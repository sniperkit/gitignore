Git Template Global/ModelSim
===

Use `git ignore add modelsim` to add this ignore template.

```
# ignore ModelSim generated files and directories (temp files and so on)
[_@]*

# ignore compilation output of ModelSim
*.mti
*.dat
*.dbs
*.psm
*.bak
*.cmp
*.jpg
*.html
*.bsf

# ignore simulation output of ModelSim
wlf*
*.wlf
*.vstf
*.ucdb
cov*/
transcript*
sc_dpiheader.h
vsim.dbg
```