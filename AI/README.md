sudo apt install python3-full
python3 -m venv ./env
source ./env/bin/activate
pip install flask
# Using a lot of data and hard drives
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple -U FlagEmbedding
git clone https://github.com/FlagOpen/FlagEmbedding.git
cd FlagEmbedding
pip install  .
