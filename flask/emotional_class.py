import numpy as np
import tensorflow as tf
class texttoemotion():
    def __init__(self,path):
        self.model = tf.keras.models.load_model(path)
    def get_emotion(self,text):
        categories = ['anger', 'fear', 'happy', 'love', 'sadness', 'surprise']
        return categories[np.argmax(self.model.predict(np.array([text.lower()]))[0])]
tte = texttoemotion('snug/mylastbraincells')
print(tte.get_emotion('I want to hurt somebody'))

