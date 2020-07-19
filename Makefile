dep-update:
	(cd ./dep/github.com/apache && ./update.sh)
	(cd ./dep/github.com/streamnative && ./update.sh)
	(cd ./dep/github.com/Shopify && ./update.sh)

.PHONY: update-dep
