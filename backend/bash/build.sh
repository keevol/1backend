BUILD_PATH=$1
AUTHOR=$2
PROJECT_NAME=$3
RECIPE_FOLDER=$4  # (eg: "go", "nodejs-whatever")
PROJECT_PATH=$5 # absolute
# envar is assumed to be in build folder under name "envars"

cp $PROJECT_PATH/tech-pack/$RECIPE_FOLDER/Dockerfile $BUILD_PATH/Dockerfile
INTERNALIP=$(ip route get 8.8.8.8 | head -1 | cut -d' ' -f8)
cd $BUILD_PATH

# image and container name is the same
CONTAINER=$AUTHOR"_"$PROJECT_NAME

sudo docker build -t $CONTAINER . || exit 1
sudo docker stop $CONTAINER
sudo docker rm $CONTAINER

sudo docker run --name $CONTAINER -p=8883 -d $CONTAINER
