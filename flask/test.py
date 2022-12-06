# import tensorflow as tf
# import numpy as np
# model = tf.keras.models.load_model('snug/aathma_is_cringe')
# print(round(model.predict(np.array(['I love you'.lower()]))[0][0]))
# # For easier use
#
#
# class SafeTrackerModel():
#     def __init__(self,path):
#         self.model =  tf.keras.models.load_model(path)
#     def predict(self,text):
#         # <0 = not  a threat
#         # >1 = threat
#         return round(self.model.predict(np.array([text.lower()]))[0][0])
#     def vectorize(self,text):
#         return self.model.layers[0](text)
#
# safetrackermodel = SafeTrackerModel('snug/aathma_is_cringe')#put the path to the model on your machine here
# print(safetrackermodel.predict('Congratulations lol have fun'))
# print(safetrackermodel.vectorize('I love you'))
# import boto3
# import jwt
# import uuid
auth_token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQxMzgxNjIsImlkIjoiYXBwIiwib3JpZ19pYXQiOjE2NjI5Mjg1NjIsInVpZCI6MTZ9.TSti46_IpkRRJb8ZjKNxGzAYCdGwYSuLe2Zyf_HU9Us"
#
# s3 = boto3.resource("s3")
# bucket = "calc.masa.space"
# # decode the jwt claims
# uid = jwt.decode(auth_token.split(" ")[1], options={"verify_signature": False}).get("uid")
# print(uid)
# response = s3.Bucket(bucket).upload_file("photos/p.png", str(uid) + "/twitter/" + str(uuid.uuid4()) + ".png")
# print(response)
import requests as r

data = {'ImageFile': "helo", 'TextContent': "helo",'AuthorTwitterID': "helo",'AuthorTwitterTag': "asd"}

response = r.post('http://localhost:8080/auth/twitter',json=data, headers={'Authorization': auth_token})
print(response.text)