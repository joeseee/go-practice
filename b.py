import os

src_dir = './data_txt'
des_dir = './data'
os.mkdir(des_dir)
for f in os.listdir(src_dir):
    cmd = 'cat %s/%s | base64 -d > %s/%s' % (src_dir, f, des_dir, f.replace('.txt', ''))
    print(cmd)
    os.system(cmd)
