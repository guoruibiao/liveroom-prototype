import com.tiger.camera.Capture;
import org.bytedeco.javacv.FrameGrabber;

public class main {
    public static void main(String[] args) throws FrameGrabber.Exception, InterruptedException, org.bytedeco.javacv.FrameRecorder.Exception {
        if(args.length != 2) {
            System.out.println("参数错误！\n 格式为：" +
                    "java jar xxx.jar rtmp-publish-address framerate");
            System.exit(-1);
        }
        String publishAddress = args[0];
        Integer frameRate = Integer.parseInt(args[1]);
        Capture capture = new Capture();
//        capture.recordCamera("rtmp://localhost:1935/stream/example",25); //推流到rtmp服务器
        capture.recordCamera(publishAddress, frameRate);
    }
}
