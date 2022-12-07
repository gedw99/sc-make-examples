# https://www.digitalocean.com/community/tutorials/how-to-set-up-a-continuous-delivery-pipeline-with-flux-on-digitalocean-kubernetes

# https://github.com/phoban01/cue-flux-controller

export GITHUB_USER=joeblew
export GITHUB_TOKEN=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



print:

dep:
	# https://fluxcd.io/docs/installation/
	gofish install flux
	echo ". <(flux completion zsh)" >> ~/.zshrc

	# https://github.com/abiosoft/colima
	brew install colima

	gofish install kubectl
	gofish install kind

pre:
	colima start --with-kubernetes 
	flux check --pre

boot:
	flux bootstrap github \
		--owner=$(GITHUB_USER) \
		--repository=flux-config \
		--branch=main \
		--path=./clusters/my-cluster \
		--personal


	

