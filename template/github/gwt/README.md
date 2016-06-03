Git Template GWT
===

Use `git ignore add gwt` to add this ignore template.

```
*.class

# Package Files #
*.jar
*.war

# gwt caches and compiled units #
war/gwt_bree/
gwt-unitCache/

# boilerplate generated classes #
.apt_generated/

# more caches and things from deploy #
war/WEB-INF/deploy/
war/WEB-INF/classes/

#compilation logs
.gwt/

#caching for already compiled files
gwt-unitCache/

#gwt junit compilation files
www-test/

#old GWT (1.5) created this dir
.gwt-tmp/
```