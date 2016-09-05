#!/usr/bin/python -u
###
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
#
# Copyright 2015 Arduino Srl (http://www.arduino.org/)
#
# authors:
# sergio@arduino.org
#
###

import argparse, requests, sys, os

_version = '0.0.1'
_serv_file = 'otafile'

def _split_file(filename, nlines):
	"""Splits the input file every nlines lines."""
	counter = 1
	file_list = []
	input_file = open(filename, 'r').read().split('\n')

	for lines in range(0, len(input_file), nlines):
		# First, get the list slice
		output_data = input_file[lines:lines + nlines]

		output_file = open(filename + "." + str(counter) + '.out', 'w')
		output_file.write('\n'.join(output_data))
		output_file.close()
		#print "Chunk #" + str(counter) + ": " + output_file.name
		file_list.append(output_file.name)
		counter += 1

	return file_list

def _upload_files(file_list, filename, address, port = 80):
	try:
		url = "http://" + address + ":" + str(port) + "/"+ _serv_file
		totchunk = len(file_list)
		numchunk = 0
		payload = {'totchunk': str(totchunk), 'numchunk': str(numchunk) }
		headers = {'Content-type': 'multipart/form-data'}
		print "Sending " + filename + " to host " + address
		for index, fl in enumerate(file_list):
		#for fl in iter(file_list):
			index += 1
			f = open(fl, 'rb')

			files = {'file': f}
			payload['numchunk'] = str(index)
			r = requests.post(url, files=files, params=payload, headers=headers, timeout=15 )
			print "["+str(index) + " / " + str(len(file_list)) + "] Done..."

			#print str(r.status_code) + " " + r.content
		print 'Upload ' + r.content
	except requests.RequestException as req_error:
		print req_error
		sys.exit(-1)

def _remove_files(file_list):
	try:
		for fl in iter(file_list):
			#print "Deleting file : " + fl
			os.remove(fl)
	except Exception as ex:
		print ex
		sys.exit(-1)

parser = argparse.ArgumentParser(prog='arduino_mcuota',
	#usage='%(prog)s ',
	prefix_chars='-',
	formatter_class=argparse.ArgumentDefaultsHelpFormatter,
	description='Arduino MCU-OTA Upload Tool - Upload sketches to the microcontroller of the board Over The Air via ESP8266 - Copyright 2016 Arduino Srl',
	epilog='')

requiredNamed = parser.add_argument_group('required named arguments')
requiredNamed.add_argument('-f', '--file', type=file, help='firmware file to upload', required=True)
requiredNamed.add_argument('-i', '--ip', help='ip of the esp8266 based board to upload to', required=True)

parser.add_argument('-p', '--port', type=int , default=80, help='network port')
parser.add_argument('-l', '--lines', type=int , default=20, help='max lines for each splitted files')
#parser.add_argument('-v', '--version', dest='lines', help='version number')

args = parser.parse_args()

_list = _split_file(args.file.name, args.lines)
_upload_files(_list, args.file.name, args.ip, args.port)
_remove_files(_list)

#sys.exit(0)
