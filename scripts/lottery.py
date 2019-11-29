import requests
import json
import smtplib
import time
from alerter import alert

while True:
	r = requests.post('https://megamillions.com/cmspages/utilservice.asmx/GetLatestDrawData')
	if r.status_code != 200:
		raise Exception(f'Bad status: {r.reason} on request for latest draw info.')
	txt = ('{' + '{'.join(r.text.split('{')[1:])).split('</string')[0]
	response = json.loads(txt)
	jackpot_info = response['Jackpot']
	if not jackpot_info['Winners'] and jackpot_info['Verified'] and jackpot_info['CurrentCashValue'] > 100000000:
		alert(f'Big boy jackpot upcoming on: {jackpot_info["PlayDate"]} for cash value of {jackpot_info["NextCashValue"]}')
	time.sleep(24 * 60 * 60)