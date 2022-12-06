from flask import Flask, request
import tweepy
import tensorflow as tf
import numpy as np
import matplotlib
import matplotlib.pyplot as plt
from flask_cors import CORS
import requests as r
from flask import jsonify, make_response
from PIL import Image
import boto3
import jwt
import uuid
import random
matplotlib.use("Agg")

def twitter_bot(enterUser,auth_token):

    class SafeTrackerModel():
        def __init__(self,path):
            self.model = tf.keras.models.load_model(path)
        def predict(self,text):
            return round(self.model.predict(np.array([text.lower()]))[0][0])

    safetrackermodel = SafeTrackerModel('model')
    

    auth = tweepy.OAuthHandler("0OxI9RXlEgK6lluxT3HKxWUw4", "4EEoCauKcpBNixlp5ulgvWAKJhU9A9tTY3f4cCKjD8LGWxWPfU")
    auth.set_access_token("1544411510394327040-rTrgobFbkX9UXZK667hKpT59qjS3eo", "Hy62AXEJA5Hos6T6QBwkm9PM2WE7UZ9Ivmx5Ac9CS7bZM")
    api = tweepy.API(auth)

    class texttoemotion():
        def __init__(self):
            self.model = tf.keras.models.load_model('mylastbraincells')
        def get_emotion(self,text):
            categories = ['Anger', 'Fear', 'Energetic', 'Friendly', 'Dark', 'Surprise']
            return categories[np.argmax(self.model.predict(np.array([text.lower()]))[0])]
    tte = texttoemotion()

    tweets = api.user_timeline(screen_name = enterUser,
                            count = 3200,
                            include_rts = False,
                            tweet_mode = 'extended'
                            )
    myint = 0 
    for info in tweets[:100]:
        myint+=1

        b = info.full_text.split()
        for item in b:
            if '@' in item or '#' in item:
                b.remove(item)
        finalB = ' '.join(b)

        arr = [
            # " I have a loaded gun and I will kill everyone I see tomorrow. "
            ""
        ]

        finalB += random.choice(arr)

        print(finalB)
        print(safetrackermodel.predict(finalB))
        print("\n\n")

        if safetrackermodel.predict(finalB) > 0:
                emotional_spectrum = {'Anger':0,'Fear':0,'Energetic':0,'Friendly':0,'Dark':0,'Surprise':0,'Neutral':0}

                for info in tweets[:myint]:
                    lol = tte.get_emotion(finalB)
                    pc = []
                    for item in list(emotional_spectrum.keys()):
                        emotional_spectrum[lol]+=1

                    for item in list(emotional_spectrum.keys()):
                        pc.append(emotional_spectrum[item])

                percentages = [(item/sum(pc))*100 for item in pc]

                labels = [f'Anger:{round(percentages[0])}',f'Fear:{round(percentages[1])}',f'Energetic:{round(percentages[2])}',f'Friendly:{round(percentages[3])}',f'Dark:{round(percentages[4])}',f'Surprise:{round(percentages[5])}',f'Neutral:{round(percentages[5])}']

                plt.cla()
                plt.pie(pc, labels = labels)

                plt.savefig('photos/p.png')

                s3 = boto3.resource("s3")
                bucket = "calc.masa.space"
                svid = str(uuid.uuid4())
                # decode the jwt claims
                uid = jwt.decode(auth_token.split(" ")[1], options={"verify_signature": False}).get("uid")
                response = s3.Bucket(bucket).upload_file("photos/p.png", str(uid) + "/twitter/" + svid + ".png")

                locationOfPieChart = "https://s3.us-west-2.amazonaws.com/" + bucket + "/" + str(uid) + "/twitter/" + svid + ".png"

                r.post('http://localhost:8080/auth/twitter',json={"ImageFile": locationOfPieChart, 'TextContent': finalB,'AuthorTwitterID':enterUser,'AuthorTwitterTag':api.get_user(screen_name = enterUser).name}, headers={'Authorization': auth_token})

app = Flask(__name__)
CORS(app)

@app.route('/',methods=["GET","POST"])
def index():
    if request.method=="POST":
        enterUser = str(request.form['TwitterScannedAccountID'])
        auth_token = str(request.headers['Authorization'])
        twitter_bot(enterUser,auth_token)

        return "ok"

    return "no"


if __name__ == "__main__":
    app.run()










