PROJ_DIR=`pwd`
VENV=${PROJ_DIR}/.env
PROJ_NAME=WordTools

if [ ! -e ${VENV} ];then
    virtualenv --never-download --prompt "(${PROJ_NAME})" ${VENV} -p $(type -p python3.7)
fi

source ${VENV}/bin/activate 

export PYTHONPATH=${PROJ_DIR}

export PROJ_NAME
export PROJ_DIR
