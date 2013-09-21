import unittest
import subprocess
import socket
import os
import time

class TestSocketAPI(unittest.TestCase):
    cmd = ['./main']
    sock_f = 'sock'

    def _unlink(self):
        try:
            os.unlink(self.sock_f)
        except FileNotFoundError:
            pass

    def setUp(self):
        self._unlink()

    def tearDown(self):
        self._unlink()

    # def createSocket(self):
    #     self.proc = subprocess.Popen(['./unmc-server', self.sock_f])
    #     self.assertTrue(os.path.exists(self.sock_f))

    #     self.socket = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    #     self.socket.connect(self.sock_f)

    def testBadCall(self):
        with self.assertRaises(subprocess.CalledProcessError) as exc:
            output = subprocess.check_output(self.cmd, stderr=subprocess.STDOUT, timeout=1)

        self.assertEqual(exc.exception.returncode, 1)
        self.assertIn(b'Usage:', exc.exception.output)

    def testCreatesSocket(self):
        proc = subprocess.Popen(self.cmd+[self.sock_f])
        time.sleep(1)
        self.assertTrue(os.path.exists(self.sock_f))
        proc.terminate()

    @unittest.expectedFailure
    def testUnlinksSocket(self):
        proc = subprocess.Popen(self.cmd+[self.sock_f])
        time.sleep(1)
        proc.terminate()
        self.assertFalse(os.path.exists(self.sock_f))

    def testWaits(self):
        proc = subprocess.Popen(self.cmd+[self.sock_f])
        time.sleep(1)
        self.assertIs(proc.poll(), None)
        proc.terminate()

    # def testVersion(self):
    #     self.createSocket()
    #     self.socket.send('VERSION')
    #     data = self.socket.recv(1024)
    #     print(data)

if __name__ == '__main__':
    unittest.main()
