package javaio.teste;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.Locale;
import java.util.Scanner;

public class TesteLeitura2 {

	public static void main(String[] args) throws FileNotFoundException {

		//FileInputStream fs = new FileInputStream("lorem.txt");
		//InputStreamReader inReader = new InputStreamReader(fs);
		//BufferedReader buffer = new BufferedReader(inReader);
		
		Scanner scanner = new Scanner(new File("contas.csv"));
		while(scanner.hasNext()) {
			String line = scanner.nextLine();
			Scanner lineScanner = new Scanner(line);
			
			lineScanner.useLocale(Locale.US);
			
			lineScanner.useDelimiter(",");

			String type = lineScanner.next();
			int numero = lineScanner.nextInt();
			int agencia = lineScanner.nextInt();
			String titular = lineScanner.next();
			double saldo = lineScanner.nextDouble();
			
			System.out.format(new Locale("pt", "BR"),"%s - %d-%d %s: %.02f %n", type, numero, agencia, titular, saldo);
		}
		
		scanner.close();
		
		
	}

}
