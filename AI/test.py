from FlagEmbedding import FlagModel
import time
import numpy as np
from flask import Flask, request, jsonify

model = FlagModel('models',
                  query_instruction_for_retrieval="Represent this sentence for searching relevant passages:",
                  use_fp16=True)
                  
def GetFeatures(words):
    embeddings = model.encode(words)
    return embeddings

app = Flask(__name__)

def get_embeddings(words):
    embeddings = model.encode(words)
    return embeddings

@app.route('/get_features', methods=['POST'])
def get_features():
    data = request.json
    words = data.get('words', [])
    embeddings = get_embeddings(words)
    for index, word in enumerate(words):
        if len(word) == 0:
             return jsonify({
            "error": "Bad Request",
            "message": "input string must not empty, please check index:%d"%index
        }), 400
    # 确保所有的numpy数组都被转换为Python列表
    if len(words) == 1:
        embeddings = [float(val) for val in embeddings.flatten()]
    else:
        embeddings = [[float(val) for val in embedding.flatten()] for embedding in embeddings]
    return jsonify({"embeddings": embeddings})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)