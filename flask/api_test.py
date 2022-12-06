from flask import Flask, request, jsonify
from PIL import Image
import requests as re

app = Flask(__name__)

@app.route("/", methods=["POST"])
def process_image():
    file = request.files['image']
    # Read the image via file.stream
    img = Image.open(file.stream)
    img.show()

    return jsonify({'msg': 'success', 'size': [img.width, img.height]})
app.run()