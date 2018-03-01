#!/usr/bin/env python
# -*- coding: utf8 -*- 

import praw
import config
import footer
import time
import urllib
from datetime import datetime

shrug = '¯\_(ツ)_/¯'
codeShrug = '`¯\_(ツ)_/¯`'
codeShrugSpace = urllib.unquote("%20%20%20%20¯\_(ツ)_/¯")
shoulders = '¯\\\_(ツ)_/¯'
decapitated = '¯\ _(ツ)_/¯'
wiseguy = '¯\_(ツ)_/¯\\'

def bot_login():
    r = praw.Reddit(username = config.username,
			password = config.password,
			client_id = config.client_id,
			client_secret = config.client_secret,
			user_agent = "u/LimbRetrieval-Bot")
    return r

def run_bot(r):
	subreddit = r.subreddit('all')
	# try:
	comments = subreddit.stream.comments()
	for comment in comments:
		text = comment.body
		author = comment.author
		run_bot.sub = str(comment.subreddit)

		if (codeShrug.decode('utf-8') in text) or (codeShrugSpace.decode('utf-8') in text):
			try:
				file = open("log.txt", "a")
			except IOError:
				file = open("log.txt", "w")

			commentTime = str(datetime.now())
			file.write(commentTime + ': ' + str(author) + ' (in r/' + run_bot.sub + ') - found in code block' '\n')
			file.close()
			break

		elif wiseguy.decode('utf-8') in text:
			try:
				file = open("log.txt", "a")
			except IOError:
				file = open("log.txt", "w")

			commentTime = str(datetime.now())
			file.write(commentTime + ': ' + str(author) + ' (in r/' + run_bot.sub + ')' '\n')
			file.close()

			comment.reply("ಠ_ಠ" + footer.footerComment + footer.explainationLink)
			break
		
		elif shrug.decode('utf-8') in text: #decode required for unicode characters
		#create/open log.txt
			try:
				file = open("log.txt", "a")
			except IOError:
				file = open("log.txt", "w")
			
			commentTime = str(datetime.now())
			file.write(commentTime + ': ' + str(author) + ' (in r/' + run_bot.sub + ')' '\n')
			file.close()

			# print "Missing limb found in: ", "(",author,")", text
			comment.reply("You dropped this \ " + footer.footerComment + footer.explainationLink)
			
		
		elif shoulders.decode('utf-8') in text:
			try:
				file = open("log.txt", "a")
			except IOError:
				file = open("log.txt", "w")

			commentTime = str(datetime.now())
			file.write(commentTime + ': ' + str(author) + ' (in r/' + run_bot.sub + ')' '\n')
			file.close()
			# print "Shoulders lost forever: ", "(",author,")", text
			comment.reply("I have retrieved these for you _ _" + footer.footerComment + footer.explainationLink)


while True:
	try:
		r = bot_login()
		run_bot(r)
		
	except Exception as ex:
		err = str(ex)
		errTime = str(datetime.now())
		file = open("log.txt", "a")
		# file = open("log.txt", "w")
		file.write(errTime + " exception caught : " + err + " in r/" + run_bot.sub + "\n") 
		file.close()
		if ("504") in err:
			time.sleep(180)
		continue
