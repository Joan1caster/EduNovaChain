from FlagEmbedding import FlagModel
import numpy as np
from flask import Flask, request, jsonify
from transformers import T5ForConditionalGeneration, T5Tokenizer


getEmbeddingsModel = FlagModel('FlagEmbedding/models',
                  query_instruction_for_retrieval="Represent this sentence for searching relevant passages:",
                  use_fp16=True)

getSummariseModel = T5ForConditionalGeneration.from_pretrained("abstract/models")
tokenizer = T5Tokenizer.from_pretrained("abstract/models")          

def GetSummarise(words, big=False):
    prefix = 'summary to zh: '
    if big:
        prefix = 'summary big to zh: '
    src_text = prefix + words
    input_ids = tokenizer(src_text, return_tensors="pt")
    generated_tokens = getSummariseModel.generate(**input_ids)
    return tokenizer.batch_decode(generated_tokens, skip_special_tokens=True)

def GetFeatures(words):
    embeddings = getEmbeddingsModel.encode(words)
    return embeddings

app = Flask(__name__)

def get_embeddings(words):
    embeddings = getEmbeddingsModel.encode(words)
    return embeddings

@app.route('/get_features', methods=['POST'])
def get_features():
    data = request.json
    words = data.get('words', [])
    for index, word in enumerate(words):
        if len(word) == 0:
             return jsonify({
            "error": "Bad Request",
            "message": "input string must not empty, please check index:%d"%index
        }), 400        
    embeddings = get_embeddings(words)
    if len(words) != 1:
        return jsonify({
            "error": "Bad Request",
            "message": "input string must be ones, please check"
        }), 400  
    else:
        embeddings = [float(val) for val in embeddings.flatten()]
    return embeddings

@app.route('/get_summary', methods=['POST'])
def get_summary():
    data = request.json
    words = data.get('words', [])
    
    for index, word in enumerate(words):
        if len(word) == 0:
             return jsonify({
            "error": "Bad Request",
            "message": "input string must not empty, please check index:%d"%index
        }), 400
    embeddings = GetSummarise(words[0])
    return jsonify({"embeddings": embeddings})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)