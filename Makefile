doc:
	rm -rf docs/_sidebar.md && ./gendoc.sh ./rust/src rs && ./gendoc.sh ./python/src py && ./gendoc.sh ./cpp/src cpp
