# https://github.com/joeblew99/gio-make-examples


print:
	## do nothign :)

	
all: fly-all minio-all

fly-all:
	cd fly-go && $(MAKE) all

minio-all:
	#cd fly-minio && $(MAKE) all



### GIT

COMMIT_MESSAGE='added more ...'
## git-commit
git-commit:
	git add --all
	git commit -am $(COMMIT_MESSAGE)

## git-pull
git-pull:
	git pull 

## git-push
git-push:
	git push 
	open $(GIT_REPO_URL)