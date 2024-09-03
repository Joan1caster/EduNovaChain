## First deploy the service need this commond, please input the commond in path/to/AI/:

# prepare python virtual env
    sudo apt install python3-full
    python3 -m venv ./env
    source ./env/bin/activate
    pip install flask

# similarity calculate
    # Will use a lot of data and hard drives
    pip install -i https://pypi.tuna.tsinghua.edu.cn/simple -U FlagEmbedding
    git clone https://github.com/FlagOpen/FlagEmbedding.git
    cd FlagEmbedding
    pip install  .

# summary generate
    git clone https://github.com/google/sentencepiece.git 
    cd sentencepiece
    mkdir build
    cd build
    cmake .. -DSPM_ENABLE_SHARED=OFF -DCMAKE_INSTALL_PREFIX=./root
    make install
    cd ../python
    python setup.py bdist_wheel
    pip install dist/sentencepiece*.whl

    ***********
    download pretrained parameters from:
    https://huggingface.co/utrobinmv/t5_summary_en_ru_zh_base_2048/tree/main
    
    move files into AI/abstract/models
    ***********

## start the service:
    source ./env/bin/activate
    python3 service.py