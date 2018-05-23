make ./bin/syz-extract
./bin/syz-extract -os ucore -arch amd64 -sourcedir ../../../../../ucore_plus/ucore/src/
make generate
make clean all
cp ./executor/syscalls_ucore.h ../../../../../ucore_plus/ucore/src/user-ucore/executor/syscalls_ucore.h
cd ../../../../../ucore_plus/ucore/
make sfsimg
# cd ../../syzkaller/src/github.com/google/syzkaller