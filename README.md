# Broadleaf Asset Generator

A simple utility to prepopulate the static asset files in the `asset.server.file.system.path` similar to what is done by `org.broadleafcommerce.cms.file.service.StaticAssetStorageServiceImpl` (without site handling). Allows us to reduce the size of the core-*.jar and the subsequent war file also.


    $ go run assetgen.go -src=~/workspaces/broadleaf/core/src/main/resources/cms/static/img -dest=/tmp
    // or
    $ go build assetgen.go
    $ ./assetgen -src=~/workspaces/broadleaf/core/src/main/resources/cms/static/img -dest=/tmp
    // for tomcat
    $ ./assetgen -src=~/workspaces/broadleaf/core/src/main/resources/cms/static/img -dest=~/apache-tomcat-8.0.0-RC3/temp

