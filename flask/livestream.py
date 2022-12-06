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
from io import BytesIO
from PIL import Image
from flask import Flask, request
from flask_cors import CORS
import base64
app= Flask(__name__)

# Load Image with PIL
def apiis(path):

    image = Image.open("/Users/evantu/documents/GoProjects/src/evantu/safetracker-server/flask/photos/a.png").convert("RGB")

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


def watch(link,auth_token,lvi):
    
    lolz= apiis(link)

    print(lolz["predictions"])

    if len(lolz['predictions']) != 0:


        img = cv2.imread(link)

        for item in lolz['predictions']:
            start_point=(int(item['x']-item['width']/2),int(item['y']-item['height']/2))

            end_point = (int(item['x']+item['width']/2),int(item['y']+item['height']/2))
            img = cv2.rectangle(img,start_point,end_point,(0,0,255),3)

        cv2.imwrite(link,img)

        #  just upload the image with boto here
        s3 = boto3.resource("s3")
        bucket = "calc.masa.space"
        svid = str(uuid.uuid4())
        # decode the jwt claims
        uid = jwt.decode(auth_token.split(" ")[1], options={"verify_signature": False}).get("uid")

        response = s3.Bucket(bucket).upload_file("/Users/evantu/documents/GoProjects/src/evantu/safetracker-server/flask/photos/a.png", str(uid) + "/livefeed/" + svid + ".png")
        # sometimes 200 sometimes fail??

        l = "https://s3.us-west-2.amazonaws.com/" + bucket + "/" + str(uid) + "/livefeed/" + svid + ".png"

        a = r.post('http://localhost:8080/auth/livethreat',json={'ImageFile': l,'LiveFeedID':lvi}, headers={'Authorization': auth_token})

        return "gun"
    return "no gun"

        
   

@app.route('/',methods=["GET","POST"])
def index():
    if request.method=="POST":
        #link = str(request.form['YouTubeLiveLink'])
        auth_token = str(request.headers['Authorization'])
        lvi = str(request.form['LiveFeedID'])
        # file = request.files['Image']
        file= request.form["Image"].replace("data:image/png;base64,","")
        file = bytes(file, "utf-8")

        with open("photos/a.png", "wb") as fh:
            fh.write(base64.decodebytes(file))

        f = watch("photos/a.png",auth_token,lvi)
        return f
    else:
        return 'Try Posting'

print("Starting app")
cors = CORS(app)
app.run(port=1111)
