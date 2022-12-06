import cv2
import numpy as np
import pafy
import youtube_dl
import requests as r
import base64
import time
import io
import uuid
import sys
import boto3
import jwt
from PIL import Image
from flask import Flask, request
from flask_cors import CORS
app= Flask(__name__)
from roboflow import Roboflow
# Load Image with PIL
def apiis(path):
    image = Image.open(path).convert("RGB")

    # Convert to JPEG Buffer
    buffered = io.BytesIO()
    image.save(buffered, quality=90, format="JPEG")

    # Base 64 Encode
    img_str = base64.b64encode(buffered.getvalue())
    img_str = img_str.decode("ascii")

    # Construct the URL
    upload_url = "".join([
        "https://detect.roboflow.com/safetracker/5",
        "?api_key=vZeO1k34YDkopMAXeu8h",
        f"&name={path}"
    ])

    # POST to the API
    re = r.post(upload_url, data=img_str, headers={
        "Content-Type": "application/x-www-form-urlencoded"
    })
    

    # Output result
    return re.json()
lolz=apiis("photos/cctv.png")

cv2.imshow("help",cv2.imread("photos/cctv.png"))

if len(lolz['predictions']) != 0:
    img = cv2.imread("photos/cctv.png")
    for item in lolz['predictions']:
        start_point=(int(item['x']-item['width']/2),int(item['y']-item['height']/2))

        end_point = (int(item['x']+item['width']/2),int(item['y']+item['height']/2))
        img = cv2.rectangle(img,start_point,end_point,(0,0,255),3)

    cv2.imwrite("photos/cctv1.jpg",img)