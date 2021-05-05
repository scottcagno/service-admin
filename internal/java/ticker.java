import java.util.Date;

class ticker {
    public static void main(String[] args) {
    System.out.println("process started successfully, running...");
    try {
            while (true) {
                System.out.println(new Date());
                Thread.sleep(60 * 1000); // sleep for one min
            }
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.printf("Hello, world!\n");
    }
}