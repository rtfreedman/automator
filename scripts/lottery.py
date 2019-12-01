import requests
import json
from alerter import alert

r = requests.post('https://megamillions.com/cmspages/utilservice.asmx/GetLatestDrawData')
if r.status_code != 200:
    raise Exception(f'Bad status: {r.reason} on request for latest draw info.')
txt = ('{' + '{'.join(r.text.split('{')[1:])).split('</string')[0]
response = json.loads(txt)
jackpot_info = response['Jackpot']
if not jackpot_info['Winners'] and jackpot_info['Verified'] and jackpot_info['CurrentCashValue'] > 100000000:
    alert(f'Big boy jackpot upcoming on: {response["NextDrawingDate"]} for cash value of {jackpot_info["NextCashValue"]}')
