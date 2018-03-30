import pandas as pd
import matplotlib.pyplot as plt

from os import listdir
from os.path import isfile, join
pathjoin = 'data/img'
for f in listdir('data/major'):
	file = join('data/major', f)
	print file
	if isfile(file):
		filename = f.split('.')[0].split('_')
		seed = filename[2]
		mom = filename[3]
		results = pd.read_csv(file)
		results = results.set_index('Time')
		
		results.plot()
		img_file = "universe_" + seed + "_" + mom + ".png"
		print img_file
		plt.savefig(join(pathjoin,img_file))
		plt.cla()
		plt.clf()
		plt.gcf().clear()
	
