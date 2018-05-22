rm /dev/shm/kafl_qemu_*
touch /dev/shm/kafl_qemu_coverage_0
touch /dev/shm/kafl_qemu_payload_0
touch /dev/shm/kafl_qemu_inpipe_0
touch /dev/shm/kafl_qemu_outpipe_0
./bin/syz-manager -config=ucore.cfg