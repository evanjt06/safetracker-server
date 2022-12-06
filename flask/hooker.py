webhook_url = 'https://discord.com/api/webhooks/1018613378744332309/NhIKxXoB1lYNvEami2Rja50_8iTP5_km88v-PPXk5q2mAVcAWqLX1A7h7ytIPTSHsCa8'
from time import sleep
from discord import Webhook, RequestsWebhookAdapter

print("starting webhook")
while True:
    webhook = Webhook.from_url(webhook_url, adapter=RequestsWebhookAdapter())
    webhook.send(content="/check")
    sleep(3600)