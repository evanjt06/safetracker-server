import discord
import numpy as np
import matplotlib.pyplot as plt
#from text2emotion import get_emotion
import tensorflow as tf
import requests as r
intents = discord.Intents(messages=True, guilds=True)
intents= discord.Intents().all()
import os

class texttoemotion():
    def __init__(self,path):
        self.model = tf.keras.models.load_model(path)
    def get_emotion(self,text):
        categories = ['Anger', 'Fear', 'Happy', 'Love', 'Sadness', 'Surprise']
        return categories[np.argmax(self.model.predict(np.array([text.lower()]))[0])]
tte = texttoemotion('mylastbraincells')
client = discord.Client(intents=intents)
token='OTkwNzU2MzU4Nzc1MjQ2ODY4.G7FSCT.AMNF9S1R1bRi8Vjbi_5l15bh1hcwpBvkzKT4cA'
@client.event
async def on_ready():
    print(f'We have logged in as {client.user}')

@client.event
async def on_message(message):
    if message.author == client.user:
        return

    if message.content.startswith('/profile'):
        content=message.content.replace('/profile ', '')
        channel = message.channel
        
        author = message.author

        messages = await channel.history(limit=500).flatten()
        await message.delete()
        msgs=[]
        idd= ''
        emotional_spectrum = {'Anger':0,'Fear':0,'Happy':0,'Love':0,'Sadness':0,'Surprise':0,'Neutral':0}
        for msg in messages:
            
           
            if msg.author.name==content:
                msgs.append(msg.content)
                idd = msg.author.id
        for mseg in msgs:
            print("MSG", mseg)
            try:
                lol = tte.get_emotion(mseg)
            except:
                lol='Neutral'
            
            emotional_spectrum[lol]+=1
        pc = []
        
        for item in list(emotional_spectrum.keys()):
            pc.append(emotional_spectrum[item])
        print("ES", pc)
        percentages = [(item/sum(pc))*100 for item in pc]
        #pc = np.array(pc)

        labels = [f'Anger:{round(percentages[0])}',f'Fear:{round(percentages[1])}',f'Conversational:{round(percentages[2])}',f'Friendly:{round(percentages[3])}',f'Harmful:{round(percentages[4])}',f'Surprise:{round(percentages[5])}',f'Neutral:{round(percentages[6])}']
        print(62, labels)

        plt.cla()
        plt.pie(pc, labels = labels)
        plt.savefig('photos/pie_chart.png')
        import tensorflow as tf
        import numpy as np
        model = tf.keras.models.load_model('model')
        cringey_lol = []
        for mseg in msgs:
            try:
                if round(model.predict(np.array([mseg.lower()]))[0][0])>=1:
                    cringey_lol.append(mseg)
            except:
                pass


        await author.send(f"Profile on {content}",file=discord.File('photos/pie_chart.png'))
        eee='\n'.join(cringey_lol)
        emb = discord.Embed(title='Messages Identified as a Threat',description=f"{eee}")
        
        await author.send(embed=emb)


    if message.content.startswith('/check'):
        channel = message.channel
        author = message.author

        messages = await channel.history(limit=500).flatten()
        await message.delete()
        msgs=[]
        
        emotional_spectrum = {'Anger':0,'Fear':0,'Happy':0,'Love':0,'Sadness':0,'Surprise':0,'Neutral':0}
        for msg in messages:
            
            msgs.append((msg.content,msg.author.name))
            try:
                lool = msg.author.roles
                x= [item.name for item in lool]
            except:
                pass
        for mseg,author in msgs:
            try:
                lol = tte.get_emotion(mseg)
            except:
                lol='Neutral'
            
            emotional_spectrum[lol]+=1
        pc = []
        
        for item in list(emotional_spectrum.keys()):
            pc.append(emotional_spectrum[item])
        percentages = [(item/sum(pc))*100 for item in pc]

        labels = [f'Anger:{round(percentages[0])}',f'Fear:{round(percentages[1])}',f'Conversational:{round(percentages[2])}',f'Friendly:{round(percentages[3])}',f'Harmful:{round(percentages[4])}',f'Surprise:{round(percentages[5])}',f'Neutral:{round(percentages[6])}']

        plt.cla()
        plt.pie(pc, labels = labels)
        plt.savefig('photos/pie_chart.png')

        import tensorflow as tf
        import numpy as np
        model = tf.keras.models.load_model('model')
        cringey_lol = []
        for mseg,author in msgs:
            try:
                if round(model.predict(np.array([mseg.lower()]))[0][0])>=1:
                    cringey_lol.append(f'{author}: {mseg}')
            except:
                pass

        await client.get_user(channel.guild.owner.id).send(f"Server Update: Profile on everybody",file=discord.File('photos/pie_chart.png'))
        eee='\n'.join(cringey_lol)
        emb = discord.Embed(title='Messages Identified as a Threat',description=f"{eee}")
        
        await client.get_user(channel.guild.owner.id).send(embed=emb)

print("starting discord bot...")
client.run(token)
