import smtplib
from email.message import EmailMessage
from getpass import getpass

pw = getpass()
def alert(s, subject="Personal Notification"):
	server = smtplib.SMTP_SSL('smtp.gmail.com', 465)
	server.login("rtfreedmannotifier", pw)
	msg = EmailMessage()
	msg.set_content(s)
	msg['Subject'] = subject
	msg['From'] = 'rtfreedmannotifier@gmail.com'
	msg['To'] = 'rtfreedman@gmail.com'
	server.send_message(msg)
	server.quit()
