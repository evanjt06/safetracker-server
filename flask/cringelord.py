import sys
from time import sleep
import webbot
from bs4 import BeautifulSoup
class EpicApi():
    def __init__(self,token):
        if token != 'ewrgnfihiuhgkj4832y78912u9ifwkjeg3yhgfhiwbgn47u89g3hv':
            sys.exit('Evan is not allow to use this cringe')
    def get_stuff(self,name):
        br = webbot.Browser()
        br.go_to(f'https://www.facebook.com/{name}')

        sleep(10)
        br.scrolly(100)
        sleep(2)
        supper = BeautifulSoup(br.get_page_source(),'lxml')
        eez = supper.find_all('span',{'class':'d2edcug0'})
        br.close_current_tab()
        return [item.get_text() for item in eez]