package main

import (
	"client/globals"
	"client/utils"
	"log" //
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("soy un log") //
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	//------------------- REALIZO CONFIGURARION MODULOS ----------------------

	globals.KernelConfig = utils.IniciarConfiguracionKernel("configKernel.json")

	if globals.KernelConfig == nil {
		log.Fatalf("No se pudo cargar la configuración kernel")
	}

	globals.CpuConfig = utils.IniciarConfiguracionCpu("configCpu.json")

	if globals.CpuConfig == nil {
		log.Fatalf("No se pudo cargar la configuración cpu")
	}

	globals.MemoryConfig = utils.IniciarConfiguracionMemory("configMemory.json")

	if globals.MemoryConfig == nil {
		log.Fatalf("No se pudo cargar la configuración memory")
	}

	globals.InterfaceConfig = utils.IniciarConfiguracionInterface("configInterface.json")

	if globals.InterfaceConfig == nil {
		log.Fatalf("No se pudo cargar la configuración Interface")
	}

	// loggeamos el valor de la config
	log.Println(globals.ClientConfig.Mensaje)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	//utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	//-----------ENVIAMOS MENSAJES ENTRE MODULOS---------------------------------

	//solicitud kernel a CPU -> respuesta CPU kernel

	utils.EnviarMensaje(globals.KernelConfig.Ip_cpu, globals.KernelConfig.Port_cpu, "soy el kernel llamando a cpu")

	//solicitud kenel a memoria -> respuesta memoria a kernel

	utils.EnviarMensaje(globals.KernelConfig.Ip_memory, globals.KernelConfig.Port_memory, "soy el kernel llamando a memoria")

	//solicitud I/O a kernel -> respuesta kerel a I/O

	utils.EnviarMensaje(globals.InterfaceConfig.Ip_kernel, globals.InterfaceConfig.Port_kernel, "soy I/O llamando a kernel")

	//solicitud I/O a memoria -> respuesta memoria a I/O

	utils.EnviarMensaje(globals.InterfaceConfig.Ip_memory, globals.InterfaceConfig.Port_memory, "soy I/O llamando a memoria")

	//solicitud CPU a memoria -> respuesta memoria a CPU*/

	utils.EnviarMensaje(globals.CpuConfig.Ip_memory, globals.CpuConfig.Port_memory, "soy CPU llamando a memoria")

	// leer de la consola el mensaje
	//utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	//utils.GenerarYEnviarPaquete()

}
