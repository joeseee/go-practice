import os

src_dir = './data'
des_dir = './data_txt'
os.mkdir(des_dir)
for f in = os.listdir(src_dir)
    cmd = 'cat %s/%s | base64 > %s/%s.txt' % (src_dir, f, des_dir, f)
    print(cmd)
    os.system(cmd)
