#!/usr/bin/env python
# coding: utf-8

# In[2]:


import pandas as pd #from sklearn.feature_extraction.text 
#import TfidfTransformer from sklearn.feature_extraction.text 
#import CountVectorizer 
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
from nltk.corpus import stopwords
import nltk
import sys
import texthero as hero
from scipy.spatial.distance import cosine
import numpy as np
np.set_printoptions(threshold=sys.maxsize)


# In[3]:


nltk.download("stopwords")


# In[4]:


from os import listdir
from os.path import isfile, join


# In[5]:


def trainForSet(directory):
    corpus = []
    listOfFiles = [f for f in listdir(directory) if isfile(join(directory, f))]
    dataset = list()
    for elem in listOfFiles:
        corpus.append(join(directory, elem)) 
        inhoudText = open(join(directory, elem),"r", encoding='utf-8') 
        dataset.append(inhoudText.read())    
        tfIdfVectorizer=TfidfVectorizer(use_idf=True, stop_words=stopwords.words('english'))    #maakt vector
        tfIdf = tfIdfVectorizer.fit_transform(dataset)  
        inhoudText.close()
        df = pd.DataFrame(tfIdf[0].T.todense(), index=tfIdfVectorizer.get_feature_names(), columns=["TF-IDF"])
        df = df.sort_values('TF-IDF', ascending=False)
        
        
       # print(tfIdf)
    return(tfIdf)


# In[6]:


def inputFileCompare(inputFile):
    #df is van de df van de inputtemplate
    setVanData= list()
    inhoudText = open(inputFile,"r", encoding='utf-8') 
    setVanData.append(inhoudText.read())    
    tfIdfVectorizer=TfidfVectorizer(use_idf=True, stop_words=stopwords.words('english'))    #maakt vector
    tfIdf = tfIdfVectorizer.fit_transform(setVanData)  
    inhoudText.close()
   
    #print(tfIdfVectorizer.get_feature_names())
    #print(hawb)
    #tfIdf = cosine_similarity(tfIdf,hawb )
    
    dfInput = pd.DataFrame(tfIdf[0].T.todense(), index=tfIdfVectorizer.get_feature_names(), columns=["TF-IDF"])
    dfInput = dfInput.sort_values('TF-IDF', ascending=False)
    print (dfInput)
    return tfIdf


# In[7]:


bol = trainForSet("input/BL") #maakt datasets aan 
seawb = trainForSet("input/SWB")
dl = trainForSet("input/DL")
hawb = trainForSet("input/HAWB")
heat = trainForSet("input/HEAT")
inv = trainForSet("input/INV")
order = trainForSet("input/ORDER")
pl = trainForSet("input/PL")
plinv = trainForSet("input/PL+INV") 

resultaat = inputFileCompare(bol,"input/INV/205153516000.txt")
print(resultaat)

output = inputFileCompare("input/HAWB/HAWB000.txt")
print(output)

# file = open("input/INV/205153516000.txt")
# data = file.read()
# file.close()

(cosine_similarity(pl, order, dense_output=True)) # methode om beide te vergelijken en er een 
# output te krijgen om te controleren  of deze 2 overeen komen

# ds = list(data)
# tfIdfVectorizer=TfidfVectorizer(use_idf=True, stop_words=stopwords.words('english'))    #maakt vector
# tfIdf = tfIdfVectorizer.fit_transform(data)  
# df = pd.DataFrame(tfIdf[0].T.todense(), index=tfIdfVectorizer.get_feature_names(), columns=["TF-IDF"])
