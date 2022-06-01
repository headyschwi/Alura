package javaio.teste;

import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Properties;

public class TesteProps {
	public static void main(String[] args) throws IOException {
		Properties props = new Properties();
		
		props.setProperty("teste", "senhaTeste");
		props.store(new FileOutputStream("conf.properties"), "algum comentário");

		Properties propsReader = new Properties();
		propsReader.load(new FileInputStream("conf.properties"));
		
		String teste = propsReader.getProperty("teste");
		System.out.println(teste);
	}
}
